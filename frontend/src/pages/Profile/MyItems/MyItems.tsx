import Icon from "components/UI/Icon";
import TextField from "components/UI/TextField";
import { IPFS_VIEW_URL } from "consts/baseUrl";
import { User } from "models/User";
import { HttpClientProviderContext } from "providers/HttpClientProvider";
import { useState, useContext, useEffect } from "react";
import { useLocation } from "react-router-dom";
import { useAccount } from "wagmi";
import { Book } from "features/book";
import styles from "./MyItems.module.scss";
import Typography from "components/UI/Typography";
import { ROUTE_TO_BUTTON_TEXT_MAPPER } from "features/user/components/ProfileSideBar/ProfileSideBar";

const MyItems = () => {
  const location = useLocation();
  const type = location.pathname.split("/")[3];
  const { address } = useAccount();
  const [user, setUser] = useState<User>();
  const { userApi } = useContext(HttpClientProviderContext);

  useEffect(() => {
    (async () => {
      const res = await userApi.GetUserWithType(address as string, type);
      setUser(res);
    })();
  }, [type]);

  return (
    <div>
      <Typography variant="h1" style={{ textAlign: "center" }}>
        {ROUTE_TO_BUTTON_TEXT_MAPPER[type]}
      </Typography>
      <div className={styles["search-bar"]}>
        <TextField
          value={""}
          onChange={() => {}}
          placeholder="Name, Author..."
          endIcon={<Icon variant="search" onClick={() => {}} />}
        />
      </div>
      <div className={styles["book-grid"]}>
        {type === "books"
          ? user?.boughtBooks?.map((item) => {
              let coverImage = IPFS_VIEW_URL + item.coverImage.substring(7);
              return <Book key={item.address} book={item} url={coverImage} small></Book>;
            })
          : user?.publishedBooks?.map((item) => {
              let coverImage = IPFS_VIEW_URL + item.coverImage.substring(7);
              return <Book key={item.address} book={item} url={coverImage} small></Book>;
            })}
      </div>
    </div>
  );
};

export default MyItems;
