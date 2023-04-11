import { Outlet } from "react-router-dom";
import { Footer, NavBar } from "features/navigation";
import { ProfileSideBar } from "features/user";
import styles from "./Main.module.scss";
import { useSelector } from "react-redux";
import { RootState } from "features/redux";
import { useAccount } from "wagmi";
import { HttpClientProviderContext } from "providers/HttpClientProvider";
import { useContext } from "react";
import { UserReqDTO } from "models/User";

const Main = () => {
  const { address, isConnected } = useAccount();
  const open = useSelector((state: RootState) => state.profileSideBar.open);
  const { userApi } = useContext(HttpClientProviderContext);

  if (isConnected) {
    userApi.CreateUser({ address } as UserReqDTO);
  }

  return (
    <>
      <> {open && isConnected && <ProfileSideBar />}</>
      <div className={styles.root}>
        <NavBar />
        <Outlet />
        <Footer />
      </div>
    </>
  );
};

export default Main;
