import { PropsWithChildren } from "react";
import { PAGE_TYPE } from "features/auth/models/AuthModel";
import { useAccount } from "wagmi";
import { Navigate } from "react-router-dom";

interface IPageNavigatorProps {
  pageType: PAGE_TYPE;
}

const PageNavigator = ({ pageType, children }: PropsWithChildren<IPageNavigatorProps>) => {
  const { isConnected } = useAccount();

  if (pageType === PAGE_TYPE.PRIVATE && !isConnected) {
    return <Navigate to="/" />;
  } else {
    return children as JSX.Element;
  }
};

export default PageNavigator;
