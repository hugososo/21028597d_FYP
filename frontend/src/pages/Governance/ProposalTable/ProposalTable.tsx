import Label from "components/UI/Label";
import Typography from "components/UI/Typography";
import { Proposal, proposalStateToText } from "models/Proposal";
import { useNavigate } from "react-router-dom";
import { formatDate } from "utils/format";
import { useProvider, useSigner } from "wagmi";
import styles from "./ProposalTable.module.scss";

type Props = React.HTMLAttributes<HTMLTableElement> & { proposals: Proposal[] };

const ProposalTable = (props: Props) => {
  const { proposals, className, ...tableProps } = props;
  const navigate = useNavigate();
  return (
    <table className={`${styles.table} ${className}`} {...tableProps}>
      <thead>
        <tr>
          <th>Proposal</th>
          <th>Agree</th>
          <th>Against</th>
          <th>Total Votes</th>
        </tr>
      </thead>
      <tbody>
        {proposals.map((item) => {
          return (
            <tr key={item.id} onClick={() => navigate(`/governance/proposal/${item.id}`, { state: { ...item } })}>
              <td>
                <Typography>{item.title}</Typography>
                <div className={`${styles["flex-row"]} ${styles.des}`}>
                  <Label color="blue1" text={proposalStateToText(item.state ?? undefined)}></Label>
                  <Typography variant="h4" className={styles["propose-date"]}>
                    {`Proposed on: ${formatDate(item.voteStart || new Date())}`}
                  </Typography>
                </div>
              </td>
              <td className={styles.agree}>{item.upVoteCount}</td>
              <td className={styles.against}>{item.downVoteCount}</td>
              <td>{Number(item.upVoteCount) + Number(item.downVoteCount)}</td>
            </tr>
          );
        })}
      </tbody>
    </table>
  );
};

export default ProposalTable;
