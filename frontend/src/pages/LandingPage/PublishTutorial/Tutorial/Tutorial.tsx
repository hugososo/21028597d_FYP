import styles from "./Tutorial.module.scss";
import Typography from "components/UI/Typography";
import { ReactComponent as Tutorial1 } from "assets/images/tutorial1.svg";
import { ReactComponent as Tutorial2 } from "assets/images/tutorial2.svg";
import { ReactComponent as Tutorial3 } from "assets/images/tutorial3.svg";

interface ITutorialProp {
  svg: "1" | "2" | "3";
  title: string;
  subtitle: string;
}

const TutorialSVG = ({ svg }: Pick<ITutorialProp, "svg">) => {
  switch (svg) {
    case "1":
      return <Tutorial1 className={styles.image} />;
    case "2":
      return <Tutorial2 className={styles.image} />;
    case "3":
      return <Tutorial3 className={styles.image} />;
    default:
      return null as never;
  }
};

const Tutorial = ({ svg, title, subtitle }: ITutorialProp) => {
  return (
    <div className={`${styles["flex-col"]} ${styles.root}`}>
      <TutorialSVG svg={svg} />
      <div className={styles["tutorial-text"]}>
        <Typography variant="h2" style={{ color: styles.secondaryBlue1, marginBottom: "20px" }}>
          {title}
        </Typography>
        <Typography variant="subtitle2">{subtitle}</Typography>
      </div>
    </div>
  );
};

export default Tutorial;
