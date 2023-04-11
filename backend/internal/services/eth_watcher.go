package services

import (
	"comp4913-backend/configs"
	"comp4913-backend/internal/contracts"
	"comp4913-backend/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"math/big"
	"net/http"
	"strings"
)

type EthWatcherService struct {
	db      *gorm.DB
	client  *ethclient.Client
	bookSvc *BookService
	ctx     context.Context
}

func NewEthWatcherService(db *gorm.DB, client *ethclient.Client, bookSvc *BookService, ctx context.Context) *EthWatcherService {
	return &EthWatcherService{db: db, client: client, bookSvc: bookSvc, ctx: ctx}
}

func NewEthClient(url string) *ethclient.Client {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (s EthWatcherService) VoteIndexing(contract string) {
	lastBlock, err := s.client.BlockNumber(s.ctx)
	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(int64(lastBlock)),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := s.client.FilterLogs(s.ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(contracts.DebookDAOMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, eventLog := range logs {
		switch eventLog.Topics[0].Hex() {
		case contracts.ProposalCreatedBytesHash.Hex():
			var data contracts.DebookDAOProposalCreated
			err := contractAbi.UnpackIntoInterface(&data, "ProposalCreated", eventLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Log Name: ProposalCreated:", common.BigToHash(data.ProposalId))
		case contracts.VotedBytesHash.Hex():
			var data contracts.DebookDAOVoted
			err := contractAbi.UnpackIntoInterface(&data, "Voted", eventLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			data.Voter = common.HexToAddress(eventLog.Topics[1].Hex())
			vote := &models.Vote{
				ProposalId: common.BigToHash(data.ProposalId).String(),
				Voter:      data.Voter.Hex(),
				Selection:  data.Selection != 0,
				Reason:     data.Reason,
			}
			s.db.Create(vote)
		}
	}
}

func (s EthWatcherService) BookIndexing(contract string) {
	startBlock := s.bookSvc.GetBookLastBlock()
	lastBlock, err := s.client.BlockNumber(s.ctx)
	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(startBlock)),
		ToBlock:   big.NewInt(int64(lastBlock)),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := s.client.FilterLogs(s.ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(contracts.ERC721FactoryMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, eventLog := range logs {
		switch eventLog.Topics[0].Hex() {
		case contracts.BookCreatedBytesHash.Hex():
			var data contracts.ERC721FactoryBookCreated
			err := contractAbi.UnpackIntoInterface(&data, "BookCreated", eventLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			url := fmt.Sprintf("%s%s", configs.IPFS_VIEW_URL, strings.TrimPrefix(data.Url, "ipfs://"))
			res := HttpGet(url)
			bookMetadata := models.BookMetadata{}
			err = json.Unmarshal(res, &bookMetadata)
			if err != nil {
				log.Fatal(err)
			}
			book := bookMetadata.MapToBook(lastBlock, data)
			s.db.FirstOrCreate(book)
		}
	}
}

func (s EthWatcherService) BookTransferIndexing() {
	//startBlock := s.bookSvc.GetBookLastBlock()
	lastBlock, _ := s.client.BlockNumber(s.ctx)
	booksAddress := s.bookSvc.GetBooksAddress()
	contractAddress := []common.Address{}
	for _, address := range booksAddress {
		contractAddress = append(contractAddress, common.HexToAddress(address))
	}
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(int64(lastBlock)),
		Addresses: contractAddress,
	}

	logs, err := s.client.FilterLogs(s.ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	for _, eventLog := range logs {
		switch eventLog.Topics[0].Hex() {
		case contracts.TransferBytesHash.Hex():
			from := common.HexToAddress(eventLog.Topics[1].Hex()).Hex()
			to := common.HexToAddress(eventLog.Topics[2].Hex()).Hex()
			tokenAddress := eventLog.Address.Hex()
			book := &models.Book{Address: tokenAddress}
			err := s.db.Clauses(clause.OnConflict{DoNothing: true}).Model(&models.User{UserAddress: from}).Association("BoughtBooks").Delete(book)
			if err != nil {
				log.Fatal(err.Error())
			}
			err = s.db.Clauses(clause.OnConflict{DoNothing: true}).Model(&models.User{UserAddress: to}).Association("BoughtBooks").Append(book)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

func (s EthWatcherService) SubscribeVote(contract string) {
	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(s.ctx, query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(contracts.DebookDAOMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case subErr := <-sub.Err():
			log.Fatal(subErr)
		case eventLog := <-logs:
			fmt.Println("Topic", eventLog.Topics)
			switch eventLog.Topics[0].Hex() {
			case contracts.ProposalCreatedBytesHash.Hex():
				var data contracts.DebookDAOProposalCreated
				err := contractAbi.UnpackIntoInterface(&data, "ProposalCreated", eventLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Log Name: ProposalCreated:", common.BigToHash(data.ProposalId))
			case contracts.VotedBytesHash.Hex():
				var data contracts.DebookDAOVoted
				err := contractAbi.UnpackIntoInterface(&data, "Voted", eventLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				data.Voter = common.HexToAddress(eventLog.Topics[1].Hex())
				vote := &models.Vote{
					ProposalId: common.BigToHash(data.ProposalId).String(),
					Voter:      data.Voter.Hex(),
					Selection:  data.Selection != 0,
					Reason:     data.Reason,
				}
				s.db.Create(vote)
			}
		}
	}
}

func (s EthWatcherService) SubscribeBook(contract string) {
	contractAddress := common.HexToAddress(contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(s.ctx, query, logs)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(contracts.ERC721FactoryMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case subErr := <-sub.Err():
			log.Fatal(subErr)
		case eventLog := <-logs:
			lastBlock, err := s.client.BlockNumber(s.ctx)
			if err != nil {
				log.Fatal(err)
			}
			switch eventLog.Topics[0].Hex() {
			case contracts.BookCreatedBytesHash.Hex():
				var data contracts.ERC721FactoryBookCreated
				err := contractAbi.UnpackIntoInterface(&data, "BookCreated", eventLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				url := fmt.Sprintf("%s%s", configs.IPFS_VIEW_URL, strings.TrimPrefix(data.Url, "ipfs://"))
				res := HttpGet(url)
				bookMetadata := models.BookMetadata{}
				err = json.Unmarshal(res, &bookMetadata)
				if err != nil {
					log.Fatal(err)
				}
				book := bookMetadata.MapToBook(lastBlock, data)
				s.db.FirstOrCreate(book)
			}
		}
	}
}

func (s EthWatcherService) SubscribeBookTransfer() {
	query := ethereum.FilterQuery{}

	logs := make(chan types.Log)
	sub, err := s.client.SubscribeFilterLogs(s.ctx, query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case subErr := <-sub.Err():
			log.Fatal(subErr)
		case eventLog := <-logs:
			if err != nil {
				log.Fatal(err)
			}
			switch eventLog.Topics[0].Hex() {
			case contracts.TransferBytesHash.Hex():
				from := common.HexToAddress(eventLog.Topics[1].Hex()).Hex()
				to := common.HexToAddress(eventLog.Topics[2].Hex()).Hex()
				tokenAddress := eventLog.Address.Hex()
				book := s.bookSvc.GetBookByAddress(tokenAddress)
				fmt.Println("from", from)
				fmt.Println("to", to)
				fmt.Println("book", book.Address)
				if book.Address != "" {
					err := s.db.Clauses(clause.OnConflict{DoNothing: true}).Model(&models.User{UserAddress: from}).Association("BoughtBooks").Delete(book)
					if err != nil {
						log.Fatal(err.Error())
					}
					err = s.db.Clauses(clause.OnConflict{DoNothing: true}).Model(&models.User{UserAddress: to}).Association("BoughtBooks").Append(book)
					if err != nil {
						log.Fatal(err.Error())
					}
				}
			}
		}
	}
}

func HttpGet(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	sitemap, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return sitemap
}
