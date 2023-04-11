import Typography from "components/UI/Typography";
import styles from "./Governance.module.scss";
import Button from "components/UI/Button";
import ProposalTable from "./ProposalTable/ProposalTable";
import TextArea from "components/UI/TextArea";
import { Suspense, useContext, useEffect, useState } from "react";
import Icon from "components/UI/Icon";
import Modal from "components/UI/Modal";
import TextField from "components/UI/TextField";
import { NotificationContext } from "features/notificaiton/providers/NotificationProvider";
import useLoreToken from "hooks/useLoreToken";
import { useAccount } from "wagmi";
import useDebookDAO from "hooks/useDebookDAO";
import { Proposal } from "models/Proposal";
import Spinner from "components/UI/Spinner";

const Governance = () => {
  const { address, isConnected } = useAccount();
  const [open, setOpen] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const { pushNotification } = useContext(NotificationContext);
  const { getProposalList, propose } = useDebookDAO({ pushNotification });
  const [proposals, setProposals] = useState<Proposal[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    (async () => {
      const data = await getProposalList();
      setProposals(data);
    })();
  }, [getProposalList]);

  return (
    <div className={styles.root}>
      <div className={`${styles["flex-col"]} ${styles["top-panel"]}`}>
        <Typography variant="h1">Governance</Typography>
        <Typography variant="body1" grey>
          Get stuff done for Debook DAO by creating an on-chain proposal.
        </Typography>
        <Button
          variant="tertiary"
          onClick={() => {
            setOpen(true);
          }}
        >
          Create Proposal
        </Button>
        <table className={styles.table}>
          <thead>
            <tr>
              <th colSpan={2}>Parameters</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>Proposal threshold</td>
              <td>2000</td>
            </tr>
            <tr>
              <td>Proposal delay</td>
              <td>1</td>
            </tr>
            <tr>
              <td>Voting Period</td>
              <td>1 week</td>
            </tr>
          </tbody>
        </table>
      </div>
      <div className={styles.line}></div>
      <div className={styles["main-content"]}>
        <ProposalTable proposals={proposals} />
      </div>
      <Modal open={open}>
        <div className={styles["modal"]}>
          <div className={styles["header"]}>
            <Typography variant="h3" className={styles.text}>
              Create Proposal
            </Typography>
            <Icon
              className={styles.cross}
              variant="cross"
              onClick={() => {
                setOpen(false);
                setTitle("");
                setDescription("");
              }}
            ></Icon>
          </div>
          <Typography variant="h5" className={`${styles.text} ${styles.id}`}>
            Proposer: {address}
          </Typography>
          <div className={styles["input-container"]}>
            <Typography style={{ marginBottom: "5px" }}>Title</Typography>
            <TextField
              onChange={(e) => setTitle(e.target.value)}
              placeholder="Enter proposal title"
              maxLength={31}
            ></TextField>
          </div>
          <div className={styles["input-container"]}>
            <Typography style={{ marginBottom: "5px" }}>Description</Typography>
            <TextArea
              onChange={(e) => setDescription(e.target.value)}
              placeholder="Enter proposal description"
            ></TextArea>
          </div>
          <div className={styles["input-container"]}>
            <Button
              variant="primary"
              long
              disabled={title === "" || description === ""}
              onClick={async () => {
                setLoading(true);
                await propose(title, description);
                setLoading(false);
                setOpen(false);
                setTitle("");
                setDescription("");
              }}
            >
              {loading ? <Spinner /> : "Create"}
            </Button>
          </div>
        </div>
      </Modal>
    </div>
  );
};

export default Governance;
