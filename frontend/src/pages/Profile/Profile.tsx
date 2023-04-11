import Typography from "components/UI/Typography";
import { Outlet } from "react-router-dom";
import { useAccount } from "wagmi";
import styles from "./Profile.module.scss";

const Profile = () => {
  const { address } = useAccount();
  const encodedString =
    "iVBORw0KGgoAAAANSUhEUgAAAQAAAAEACAIAAADTED8xAAAEZ0lEQVR4nOzcTasQZBqHcc/MmUGGGWHEYRgnxU1SBGUFLQoKInoRBGkVSJEV6aqFgRWCRNTCioraVEZFQbloYdkLFYSLjlEtJAiE0loFKaIkbtKSPsUNwfX7fYD/s7q4d8/iFftuXTZp3dVrRvd3bvhidP/0hmOj++fe3jK6/9njS6P7e28/O7p/6berRvf/MroOf3ICIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBEDawj8OrRh94LnjD4/uL//h8Oj+jqXZ/++/27V/dH//v7aP7v//wbtH9zfvOD+67wKQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApC1c++PO0QfeOvvG6P4Lx9eP7q/92/9G9w++tHp0f92J3aP7J345Mrq/Z8cdo/suAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBEDa4rnXV40+sOfYO6P7r/76/Oj+wctOje6vOf/56P7ejU+N7v/3sctH98+cfHl03wUgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASBMAaQIgTQCkCYA0AZAmANIEQJoASFt4dsO20QfW7147uv/vAytG919Zfcvo/pM3bBndP/PmbaP7r/1n4+j+M0tfju67AKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpi/ev3Dz6wM8Xto7ubzrwz9H9vY/8fXT/ofeuGt2/954XR/dXXjg6un/nNReN7rsApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGmLd1150+gDWx/YPrp/+vtHR/ePbLtxdP+rTy4e3d93as/o/m/3XTe6f/LoT6P7LgBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRAmgBIEwBpAiBNAKQJgDQBkCYA0gRA2sLhD5dGH/h9zTej+x9/vWt0/+Z3N43uf3po+ej+Je8/Pbr/148+GN1f9sT1o/MuAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBECaAEgTAGkCIE0ApAmANAGQJgDSBEDaHwEAAP//3AxkRicFc0kAAAAASUVORK5CYII=";

  return (
    <div className={styles.root}>
      <div className={`${styles["flex-col"]} ${styles["top-panel"]}`}>
        <div className={styles.img}>
          <img src={`data:image/png;base64,${encodedString}`} alt="Red dot" />
        </div>
        <Typography variant="h1">Profile</Typography>
        <Typography variant="body1" grey>
          {address}
        </Typography>
      </div>
      <div className={styles.line}></div>
      <div className={styles["main-content"]}>
        <Outlet />
      </div>
    </div>
  );
};

export default Profile;
