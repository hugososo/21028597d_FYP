import styles from "./BookDetail.module.scss";
import Typography from "components/UI/Typography";
import Button from "components/UI/Button";
import Icon from "components/UI/Icon";
import { Book } from "models/Book";
import { HttpClientProviderContext } from "providers/HttpClientProvider";
import { useState, useContext, useEffect } from "react";
import { useLocation } from "react-router-dom";
import { IPFS_VIEW_URL } from "consts/baseUrl";
import DataBlock from "components/DataBlock/DataBlock";
import Spinner from "components/UI/Spinner";
import useERC721 from "hooks/useERC721";
import { NotificationContext } from "features/notificaiton/providers/NotificationProvider";
import { BigNumber } from "ethers";
import { weiToEth } from "utils/format";
import useLoreToken from "hooks/useLoreToken";
import { useAccount } from "wagmi";
import { catchError } from "utils/error";

const BookDetail = () => {
  const location = useLocation();
  const address = location.pathname.split("/")[2];
  const { address: owner } = useAccount();
  const [book, setBook] = useState<Book>();
  const [price, setPrice] = useState<BigNumber>(BigNumber.from(0));
  const { bookApi } = useContext(HttpClientProviderContext);
  const coverImage = IPFS_VIEW_URL + book?.coverImage.substring(7);
  const { pushNotification } = useContext(NotificationContext);
  const { mint, cost } = useERC721({ contractAddress: address, pushNotification });
  const { useAllowance, approve, getBalanceOf } = useLoreToken({ pushNotification });
  const { data: allowance } = useAllowance(owner, address);
  const [balanceOf, setBalanceOf] = useState<string | undefined>(undefined);

  useEffect(() => {
    (async () => {
      const res = await bookApi.GetBookByAddress(address);
      setBook(res);
    })();

    (async () => {
      const res = await cost();
      setPrice(res || BigNumber.from(0));
    })();

    (async () => {
      const res = await getBalanceOf(owner);
      setBalanceOf(res || "0");
    })();
  }, []);

  const handleMint = async () => {
    if (BigNumber.from(balanceOf).lt(price)) {
      pushNotification({
        type: "warning",
        title: "You don't have enough $LORE",
        content: "Please go to Uniswap/ other DeFi to buy more $LORE",
      });
    } else if (BigNumber.from(allowance).lt(price)) {
      try {
        await approve(address);
        await mint();
      } catch (err) {
        pushNotification(catchError(err));
      }
    } else {
      await mint();
    }
  };

  if (book) {
    return (
      <div className={styles.root}>
        <div className={styles["book-detail-container"]}>
          <div className={styles.img}>
            <img src={coverImage} alt="Book" />
          </div>
          <div className={styles["book-info"]}>
            <Typography variant="h4">Author - {book?.author}</Typography>
            <Typography variant="h5">From - {book?.authorCity}</Typography>
            <Typography variant="h1">{book?.bookName}</Typography>
            <div className={styles["data-block"]}>
              <DataBlock icon="book" data={book?.bookGenre}>
                Genre
              </DataBlock>
              <DataBlock icon="book" data={book?.bookLanguage}>
                Language
              </DataBlock>
              <DataBlock icon="book" data={book?.edition}>
                Edition
              </DataBlock>
              <DataBlock icon="book" data={book?.keywords.join()}>
                Keywords
              </DataBlock>
            </div>
            <div
              className={styles.des}
              dangerouslySetInnerHTML={{ __html: book.description.replace(/\n/g, "<br />") }}
            />
            <div className={styles.price}>
              <div className={styles["flex-row"]}>
                <Icon variant="ether"></Icon>
                <Typography variant="h4">{weiToEth(price.toString())}</Typography>
              </div>
              <Button long onClick={handleMint}>
                Buy
              </Button>
            </div>
          </div>
        </div>
      </div>
    );
  } else {
    return (
      <div className={styles.root}>
        <Spinner />
      </div>
    );
  }
};

export default BookDetail;
