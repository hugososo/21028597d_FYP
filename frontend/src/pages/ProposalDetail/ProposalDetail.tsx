import ImageRenderer from "components/UI/ImageRenderer";
import styles from "./ProposalDetail.module.scss";
import Typography from "components/UI/Typography";
import Button from "components/UI/Button";
import Label from "components/UI/Label";
import Modal from "components/UI/Modal";
import RadioGroup from "components/UI/RadioGroup";
import { useContext, useState } from "react";
import TextArea from "components/UI/TextArea";
import Icon from "components/UI/Icon";
import { useLocation } from "react-router-dom";
import { Proposal } from "models/Proposal";
import { formatDate } from "utils/format";
import useDebookDAO from "hooks/useDebookDAO";
import { NotificationContext } from "features/notificaiton/providers/NotificationProvider";
import { useAccount, useContractEvent } from "wagmi";
import DAOContactABI from "contracts/abi/DebookDAO.json";
import { DebookDAOAddress } from "contracts/address";

const ProposalDetail = () => {
  const { address, isConnected } = useAccount();
  const [vote, setVote] = useState<string | undefined>(undefined);
  const [open, setOpen] = useState(false);
  const [reason, setReason] = useState("");
  const location = useLocation();
  const proposal: Proposal = location.state;
  const { pushNotification } = useContext(NotificationContext);
  const { useSignVote } = useDebookDAO({ pushNotification });
  const { sign, data, isError, isLoading, isSuccess, signTypedData, reset } = useSignVote({
    proposalId: proposal.id,
    voter: address ?? "0x0",
    selection: Number(vote),
    reason: reason,
  });

  useContractEvent({
    address: DebookDAOAddress,
    abi: DAOContactABI,
    eventName: "Console",
    listener(title, message) {
      console.log(title, message);
    },
  });

  return (
    <div className={styles.root}>
      <div className={`${styles["flex-col"]} ${styles["top-panel"]}`}>
        <Typography variant="body2" className={styles.status}>
          <Label color="blue1" text={"ACTIVE"}></Label>
          {`End Time: ${formatDate(proposal.voteEnd!)}`}
        </Typography>
        <Typography variant="h1" className={styles.title}>
          {`Proposal Title ${proposal.title}`}
        </Typography>
        <Typography variant="body1" grey>
          {`ID: ${proposal.id}· Proposed by: ${proposal.proposer}`}
        </Typography>
        <Button className={styles.vote} variant="primary" onClick={() => setOpen(true)}>
          Vote
        </Button>
      </div>
      <div className={styles.line}></div>
      <div className={styles["main-content"]}>
        <Typography variant="body1" grey className={styles.description}>
          {proposal.description}
        </Typography>
      </div>
      <Modal open={open}>
        <div className={styles["modal"]}>
          <div className={styles["header"]}>
            <Typography variant="h3" className={styles.text}>
              Voting
            </Typography>
            <Icon
              className={styles.cross}
              variant="cross"
              onClick={() => {
                setOpen(false);
                setVote(undefined);
                setReason("");
                reset();
              }}
            ></Icon>
          </div>
          <Typography variant="h5" className={`${styles.text} ${styles.id}`}>
            {`Proposal ID: ${proposal.id}`}
          </Typography>
          {isConnected ? (
            <>
              <div className={styles["input-container"]}>
                <Typography>Choose your vote</Typography>
                <RadioGroup
                  selected={vote}
                  groupName={"vote"}
                  labelValueList={[
                    { text: "Agree", value: "1", color: "green" },
                    { text: "Against", value: "0", color: "red" },
                  ]}
                  onChange={(e) => setVote(e.target.value)}
                ></RadioGroup>
              </div>
              <div className={styles["input-container"]}>
                <Typography style={{ marginBottom: "5px" }}>Reason</Typography>
                <TextArea
                  value={reason}
                  onChange={(e) => setReason(e.target.value)}
                  placeholder="Describe your Reason"
                ></TextArea>
              </div>
              <div className={styles["input-container"]}>
                <Button variant="primary" long disabled={vote === undefined} onClick={() => signTypedData()}>
                  Sign and Submit
                </Button>
              </div>
            </>
          ) : (
            <div>You need to connect your wallet first</div>
          )}
        </div>
      </Modal>
    </div>
  );
};

export default ProposalDetail;
