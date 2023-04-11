import styles from "./Banner.module.scss";
import Button from "components/UI/Button";
import Typography from "components/UI/Typography";
import book1 from "assets/images/book1.jpg";
import { Book as BookBox } from "features/book";
import { Book } from "models/Book";
import { NavigateFunction } from "react-router-dom";

interface IBannerProp {
  navigate: NavigateFunction;
}

const Banner = ({ navigate }: IBannerProp) => {
  return (
    <div className={`${styles["flex-row"]} ${styles["banner-container"]} ${styles.container}`}>
      <div className={styles["left-container"]}>
        <Typography variant="title1">
          Publish Books, <strong>EASY</strong>
        </Typography>
        <Typography variant="subtitle1">98% loyalties return to authors</Typography>
        <div className={`${styles["flex-row"]} ${styles["button-container"]}`}>
          <Button big onClick={() => navigate("/explore")}>
            Explore
          </Button>
          <Button big variant="tertiary" onClick={() => navigate("/publish")}>
            Publish
          </Button>
        </div>
      </div>
      <div className={styles["right-container"]}>
        <BookBox
          book={
            {
              bookName: "Introduction of Algorithm",
              author: "Author",
            } as unknown as Book
          }
          url={book1}
        />
      </div>
    </div>
  );
};

export default Banner;
