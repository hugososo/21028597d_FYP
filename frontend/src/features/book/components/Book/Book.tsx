import styles from "./Book.module.scss";
import ImageRenderer from "components/UI/ImageRenderer";
import Typography from "components/UI/Typography";
import Icon from "components/UI/Icon";
import { useNavigate } from "react-router-dom";
import { Book as BookType } from "models/Book";
import { NotificationContext } from "features/notificaiton/providers/NotificationProvider";
import useERC721 from "hooks/useERC721";
import { useContext, useEffect, useState } from "react";
import { BigNumber } from "ethers";
import { weiToEth } from "utils/format";

interface IBookProps {
  book: BookType;
  url: string;
  small?: boolean;
  className?: string;
}

const Book = ({ book, url, small, className }: IBookProps) => {
  const navigate = useNavigate();
  const { pushNotification } = useContext(NotificationContext);
  const [price, setPrice] = useState<BigNumber>(BigNumber.from(0));
  const { cost } = useERC721({ contractAddress: book.address, pushNotification });

  useEffect(() => {
    (async () => {
      const res = await cost();
      setPrice(res || BigNumber.from(0));
    })();
  }, []);

  return (
    <div className={`${styles.root} ${className || ""}`} onClick={() => navigate(`/explore/${book.address}`)}>
      <ImageRenderer url={url} className={`${styles.image} ${small ? styles.small : styles.big}`} />
      <div className={`${styles["book-info"]} ${small ? styles.small : styles.big}`}>
        {!small && <Typography variant="h3">{book.bookName}</Typography>}
        {small && (
          <>
            <Typography variant="h6" grey>
              {book.author}
            </Typography>
            <Typography variant="h6">{book.bookName}</Typography>
            <div className={styles["flex-row"]}>
              <Icon variant="ether"></Icon>
              <Typography variant="h5">{weiToEth(price.toString())}</Typography>
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default Book;
