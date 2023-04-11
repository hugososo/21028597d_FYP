import Typography from "components/UI/Typography";
import { Book } from "models/Book";
import styles from "./Explore.module.scss";
import TextField from "components/UI/TextField";
import Icon from "components/UI/Icon";
import { HttpClientProviderContext } from "providers/HttpClientProvider";
import { useContext, useEffect, useState } from "react";
import { Book as BookBox } from "features/book";
import { IPFS_VIEW_URL } from "consts/baseUrl";

const Explore = () => {
  const [books, setBooks] = useState<Book[]>([]);
  const { bookApi } = useContext(HttpClientProviderContext);

  useEffect(() => {
    (async () => {
      const res = await bookApi.GetBooks();
      setBooks(res);
    })();
  }, []);

  return (
    <div className={styles.root}>
      <div className={`${styles["flex-col"]} ${styles["top-panel"]}`}>
        <Typography variant="h1">Explore</Typography>
        <Typography variant="body1" grey>
          Buy any book you wants!
        </Typography>
      </div>
      <div className={styles.line}></div>
      <div className={styles["main-content"]}>
        <div className={styles["search-bar"]}>
          <TextField
            value={""}
            onChange={() => {}}
            placeholder="Name, Author..."
            endIcon={<Icon variant="search" onClick={() => {}} />}
          />
        </div>
        <div className={styles["book-grid"]}>
          {books.map((item) => {
            let coverImage = IPFS_VIEW_URL + item.coverImage.substring(7);
            return <BookBox key={item.address} book={item} url={coverImage} small></BookBox>;
          })}
        </div>
      </div>
    </div>
  );
};

export default Explore;
