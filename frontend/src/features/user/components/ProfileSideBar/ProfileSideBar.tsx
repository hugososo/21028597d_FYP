import { FC, memo } from "react";
import { useNavigate, useLocation, Location } from "react-router-dom";
import styles from "./ProfileSideBar.module.scss";
import Typography from "components/UI/Typography";
import { useAccount, useDisconnect } from "wagmi";
import Icon from "components/UI/Icon";
import { close } from "./ProfileSideBarSlice";
import { useDispatch } from "react-redux";
import PureButton from "components/UI/Button";

export enum SETTING_SECTION {
  BOOKS = "books",
  PUBLISHES = "publishes",
  ACTIVITIES = "activites",
}

export const ROUTE_TO_BUTTON_TEXT_MAPPER: { [key: string]: string } = {
  [SETTING_SECTION.BOOKS]: "My Books",
  [SETTING_SECTION.PUBLISHES]: "My Publishes",
  [SETTING_SECTION.ACTIVITIES]: "Activities",
};

const RealSideBar: FC = memo(() => {
  const navigate = useNavigate();
  const { address, isConnected } = useAccount();
  const dispatch = useDispatch();
  const { disconnect } = useDisconnect();

  return (
    <div className={styles["side-bar"]}>
      <Icon
        variant="cross"
        className={styles["close-btn"]}
        onClick={() => {
          dispatch(close());
        }}
      />
      <div className={styles["info"]}>
        <Typography variant="h3" className={styles["signed-in-as"]}>
          Connected as {address?.substring(0, 7) + "..." + address?.substring(38)}
        </Typography>
        <PureButton
          onClick={() => {
            disconnect();
            dispatch(close());
          }}
        >
          Disconnect
        </PureButton>
      </div>
      <Typography variant="h3" className={styles["category"]}>
        Settings
      </Typography>
      <Button
        onButtonClick={() => {
          navigate(`/profile/${address}/${SETTING_SECTION.BOOKS}`);
        }}
        routeName={SETTING_SECTION.BOOKS}
      />
      <Button
        onButtonClick={() => {
          navigate(`/profile/${address}/${SETTING_SECTION.PUBLISHES}`);
        }}
        routeName={SETTING_SECTION.PUBLISHES}
      />
      <Button
        onButtonClick={() => {
          navigate(`/profile/${address}/${SETTING_SECTION.ACTIVITIES}`);
        }}
        routeName={SETTING_SECTION.ACTIVITIES}
      />
    </div>
  );
});

const Button = ({
  onButtonClick,
  routeName,
}: {
  onButtonClick: (selected: SETTING_SECTION) => void;
  routeName: SETTING_SECTION;
}) => {
  return (
    <div onClick={() => onButtonClick(routeName)} className={`${styles["btn-wrapper"]} ${styles["interval"]}`}>
      <div className={styles["btn"]}>{ROUTE_TO_BUTTON_TEXT_MAPPER[routeName]}</div>
    </div>
  );
};

const SettingSideBar: FC = () => (
  <div className={styles["root"]}>
    <RealSideBar />
  </div>
);

export default SettingSideBar;
