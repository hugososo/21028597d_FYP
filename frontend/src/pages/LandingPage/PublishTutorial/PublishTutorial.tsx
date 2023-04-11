import styles from "./PublishTutorial.module.scss";
import Tutorial from "./Tutorial/Tutorial";

import { NavigateFunction } from "react-router-dom";
import Button from "components/UI/Button";

interface IPublishTutorialProp {
  navigate: NavigateFunction;
}

const PublishTutorial = ({ navigate }: IPublishTutorialProp) => {
  return (
    <div className={styles.root}>
      <div className={styles["publish-tutorial-container"]}>
        <div className={`${styles.container} ${styles["item-container"]}`}>
          <Tutorial
            svg="1"
            title={"Prepare your book and $NEAR"}
            subtitle={
              "It starts with a simple prepare of your book pdf or epub and cover files, also the $LORE coin for book publishing"
            }
          />
          <Tutorial
            svg="2"
            title={"Input book information"}
            subtitle={
              "It takes a few action to input all your book informaiton, includes book name, author, description, and the book cover."
            }
          />
          <Tutorial
            svg="3"
            title={"Publish Your Book"}
            subtitle={
              "After everything is ready, just click the confirm button, pay the gas fee, and your book will be published"
            }
          />
        </div>
        <Button big variant="tertiary" onClick={() => navigate("/publish")}>
          Go To Publish
        </Button>
      </div>
    </div>
  );
};

export default PublishTutorial;
