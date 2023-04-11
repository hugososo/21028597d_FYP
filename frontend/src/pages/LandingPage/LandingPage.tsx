import styles from "./LandingPage.module.scss";
import { useNavigate } from "react-router-dom";
import Banner from "./Banner/Banner";
import PublishTutorial from "./PublishTutorial/PublishTutorial";
import StakeInfo from "./StakeInfo/StakeInfo";

const LandingPage = () => {
  const navigate = useNavigate();

  return (
    <div>
      <Banner navigate={navigate} />
      <PublishTutorial navigate={navigate} />
      <StakeInfo navigate={navigate} />
    </div>
  );
};

export default LandingPage;
