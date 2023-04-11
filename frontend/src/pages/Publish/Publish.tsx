import Button from "components/UI/Button";
import Spinner from "components/UI/Spinner";
import TextArea from "components/UI/TextArea";
import TextField from "components/UI/TextField";
import Typography from "components/UI/Typography";
import { NotificationContext } from "features/notificaiton/providers/NotificationProvider";
import useERC721Factory from "hooks/useERC721Factory";
import { HttpClientProviderContext } from "providers/HttpClientProvider";
import React, { useContext, useState } from "react";
import { useAccount } from "wagmi";
import styles from "./Publish.module.scss";

interface IFieldProp {
  name: string;
  children: React.ReactNode;
}

const Field = ({ name, children }: IFieldProp) => {
  return (
    <>
      <div className={styles["field"]}>
        <div className={styles.name}>
          <Typography variant="h5">{name}</Typography>
        </div>
        <div>{children}</div>
      </div>
      <div className={styles.line}></div>
    </>
  );
};

const Publish = () => {
  const [price, setPrice] = useState(0);
  const { bookApi } = useContext(HttpClientProviderContext);
  const { address } = useAccount();
  const [isLoading, setIsLoading] = useState(false);
  const { pushNotification } = useContext(NotificationContext);
  const { deployERC721 } = useERC721Factory({ pushNotification });

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    // Prevent the browser from reloading the page
    e.preventDefault();
    // Read the form data
    setIsLoading(true);
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);
    formData.append("owner", address as string);
    const res = await bookApi.UploadBook(formData);
    let name = formData.get("bookName") as string;
    let symbol = formData.get("symbol") as string;
    let cost = Number(formData.get("price")?.toString());
    let ipfsUrl = "ipfs://" + res;
    await deployERC721({ name, symbol, ipfsUrl, cost });
    setIsLoading(false);
  };

  return (
    <div className={styles.root}>
      <div className={`${styles["flex-col"]} ${styles["top-panel"]}`}>
        <Typography variant="h1">Publish</Typography>
        <Typography variant="body1" grey>
          Fill in the required fields, upload the book pdf or epub and the cover image.
        </Typography>
      </div>
      <div className={styles.line}></div>
      <div className={styles["main-content"]}>
        <form method="post" onSubmit={handleSubmit}>
          <Field name="Book Name">
            <TextField name="bookName" required></TextField>
          </Field>
          <Field name="Book Name Short Form">
            <TextField name="symbol" label="Should not contains space and symbols" required></TextField>
          </Field>
          <Field name="Book Description">
            <TextArea name="description" required height={styles.area}></TextArea>
          </Field>
          <Field name="Author">
            <TextField name="author" required></TextField>
          </Field>
          <Field name="Author City">
            <TextField name="authorCity" required></TextField>
          </Field>
          <Field name="Book Genre">
            <TextField name="bookGenre" required></TextField>
          </Field>
          <Field name="Language">
            <TextField name="bookLanguage" required></TextField>
          </Field>
          <Field name="Edition">
            <TextField name="edition" required></TextField>
          </Field>
          <Field name="Keywords">
            <TextField name="keywords" label="Keywords will be split by ','"></TextField>
          </Field>
          <Field name="List Price">
            <TextField name="price" required pattern="^\d*(\.\d{0,3})?$"></TextField>
          </Field>
          <Field name="Cover Image">
            <TextField
              name="coverImage"
              required
              type="file"
              accept="image/png, image/jpeg, image/jpg, image/gif"
            ></TextField>
          </Field>
          <Field name="Book pdf/ePub/txt">
            <TextField
              name="sourceFile"
              required
              type="file"
              accept="application/pdf, application/epub+zip, text/plain"
            ></TextField>
          </Field>
          <Button className={styles.btn} type={"submit"}>
            {!isLoading ? "Confirm" : <Spinner />}
          </Button>
        </form>
      </div>
    </div>
  );
};

export default Publish;
