import { useCallback, memo } from "react";
import { useLocation, Location, matchPath, useNavigate } from "react-router-dom";
import styles from "./navBar.module.scss";
import { PAGES } from "features/navigation/models/pages";
import Icon from "components/UI/Icon";
import Typography from "components/UI/Typography";
import { useAccount, useConnect, useFeeData } from "wagmi";
import { InjectedConnector } from "wagmi/connectors/injected";
import { open } from "features/user/components/ProfileSideBar/ProfileSideBarSlice";
import { useDispatch } from "react-redux";
import Loading from "components/UI/Loading";

const matchNavPath = (path: string, nav: string) => matchPath(`/${nav}/*`, path);

const NavBar = memo(() => {
  const location = useLocation();
  const navigate = useNavigate();
  const { isConnected } = useAccount();
  const { connect } = useConnect({
    connector: new InjectedConnector(),
  });
  const onSectionButtonSelect = useCallback(
    (selected: PAGES) => {
      navigate(`/${selected}`);
    },
    [navigate]
  );

  const dispatch = useDispatch();
  const { data, isLoading } = useFeeData();

  return (
    <div className={styles["root"]}>
      <div className={styles["nav-container"]}>
        <div className={styles["left-container"]}>
          <Icon variant="logo" className={styles.logo} onClick={() => onSectionButtonSelect(PAGES.LANDING)} />
          <SectionButton onSectionButtonSelect={onSectionButtonSelect} location={location} routeName={PAGES.PUBLISH} />
          <SectionButton onSectionButtonSelect={onSectionButtonSelect} location={location} routeName={PAGES.EXPLORE} />
          <SectionButton onSectionButtonSelect={onSectionButtonSelect} location={location} routeName={PAGES.REWARDS} />
          <SectionButton onSectionButtonSelect={onSectionButtonSelect} location={location} routeName={PAGES.GOV} />
        </div>
        <div className={styles["right-container"]}>
          <div className={`${styles["flex-row"]} ${styles.center}`}>
            <Icon variant="gas" className={styles["nav-icon"]} />
            <Typography variant="h3" className={styles["nav-icon"]}>
              {isLoading ? <Loading /> : Number(data?.gasPrice) / 1e9}
            </Typography>
          </div>
          <Icon
            variant="wallet"
            className={`${styles["nav-icon"]} ${styles["btn"]}`}
            onClick={() => {
              if (isConnected) dispatch(open());
              else connect();
            }}
          />
        </div>
      </div>
    </div>
  );
});

const SectionButton = ({
  onSectionButtonSelect,
  location,
  routeName,
}: {
  onSectionButtonSelect: (selected: PAGES) => void;
  location: Location;
  routeName: PAGES;
}) => (
  <div
    onClick={() => onSectionButtonSelect(routeName)}
    className={`${styles["btn"]} ${matchNavPath(location.pathname, routeName) ? styles["selected"] : ""}`}
  >
    {routeName}
  </div>
);

export default NavBar;
