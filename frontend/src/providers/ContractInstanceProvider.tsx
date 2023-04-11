import loreTokenABI from "contracts/abi/LoreToken.json";
import stakePoolABI from "contracts/abi/StakePool.json";
import debookDAOABI from "contracts/abi/DebookDAO.json";
import erc721FactoryABI from "contracts/abi/ERC721Factory.json";
import { LoreTokenAddress, StakePoolAddress, DebookDAOAddress, ERC721FactoryAddress } from "contracts/address";
import { Contract } from "ethers";
import { createContext, PropsWithChildren } from "react";
import { useContract, useProvider, useSigner } from "wagmi";

type Props = {
  loreTokenContractWrite?: Contract | null;
  loreTokenContractRead?: Contract | null;
  stakePoolContract?: Contract | null;
  DAOContractWrite?: Contract | null;
  DAOContractRead?: Contract | null;
  ERC721FactoryContractWrite?: Contract | null;
  ERC721FactoryContractRead?: Contract | null;
};

type ContextProps = Required<Props>;

const defaultValues = {
  loreTokenContractWrite: null,
  loreTokenContractRead: null,
  stakePoolContract: null,
  DAOContractWrite: null,
  DAOContractRead: null,
  ERC721FactoryContractWrite: null,
  ERC721FactoryContractRead: null,
};

export const ContractInstanceContext = createContext<ContextProps>(defaultValues);

const ContractInstanceProvider = ({ children }: PropsWithChildren<Props>) => {
  const { data: signer } = useSigner();
  const provider = useProvider();

  const loreTokenContractWrite = useContract({
    address: LoreTokenAddress,
    abi: loreTokenABI,
    signerOrProvider: signer,
  });

  const loreTokenContractRead = useContract({
    address: LoreTokenAddress,
    abi: loreTokenABI,
    signerOrProvider: provider,
  });

  const stakePoolContract = useContract({
    address: StakePoolAddress,
    abi: stakePoolABI,
    signerOrProvider: signer,
  });

  const DAOContractRead = useContract({
    address: DebookDAOAddress,
    abi: debookDAOABI,
    signerOrProvider: provider,
  });

  const DAOContractWrite = useContract({
    address: DebookDAOAddress,
    abi: debookDAOABI,
    signerOrProvider: signer,
  });

  const ERC721FactoryContractWrite = useContract({
    address: ERC721FactoryAddress,
    abi: erc721FactoryABI,
    signerOrProvider: signer,
  });

  const ERC721FactoryContractRead = useContract({
    address: ERC721FactoryAddress,
    abi: erc721FactoryABI,
    signerOrProvider: provider,
  });

  return (
    <ContractInstanceContext.Provider
      value={{
        loreTokenContractWrite,
        loreTokenContractRead,
        stakePoolContract,
        DAOContractWrite,
        DAOContractRead,
        ERC721FactoryContractWrite,
        ERC721FactoryContractRead,
      }}
    >
      {children}
    </ContractInstanceContext.Provider>
  );
};

export default ContractInstanceProvider;
