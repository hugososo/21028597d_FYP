import styles from "./Footer.module.scss";
import Typography from "components/UI/Typography";

const Footer = () => {
  return (
    <div className={styles.root}>
      <Typography variant="body1" grey>
        @Copyright 2022 DeBook
      </Typography>
      <ul className={styles["footer-nav"]}>
        <li>
          <Typography variant="body1">Buy $LORE</Typography>
        </li>
        <li>
          <Typography variant="body1">About</Typography>
        </li>
        <li>
          <Typography variant="body1">Terms of Service</Typography>
        </li>
      </ul>
    </div>
  );
};

export default Footer;
