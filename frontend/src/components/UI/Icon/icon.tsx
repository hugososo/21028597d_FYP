import styles from "./icon.module.scss";
import { IconVariant } from "./icon.type";
import {
  AiOutlineInfoCircle,
  AiOutlineWarning,
  AiOutlineShoppingCart,
  AiOutlineSearch,
  AiOutlineUpload,
  AiOutlineSetting,
  AiOutlineControl,
} from "react-icons/ai";
import { BiWallet, BiFilterAlt, BiGasPump } from "react-icons/bi";
import { HiOutlineCheckCircle, HiOutlineExternalLink } from "react-icons/hi";
import { FaEthereum, FaExchangeAlt } from "react-icons/fa";
import { RxCross2 } from "react-icons/rx";
import { GiBookCover } from "react-icons/gi";
import { BsCheck2, BsArrowLeft, BsArrowRight } from "react-icons/bs";
import { ReactComponent as LogoIcon } from "assets/icons/icon_logo.svg";

export interface IIconProps {
  variant: IconVariant;
  bgColor?: string;
  className?: string;
  onClick?: () => void;
}

/**
 * Component for displaying icon
 * @param variant {IconVariant} type of icon
 * @param value {string} value to display on icon
 * @param color {string} background color of text icon
 * @param tooltip {string} show tooltip text
 * @param className {string} injectable className for positioning
 * @returns Icon component
 *
 * @component
 * @example
 * return (
 *  <Icon variant="text" value="F" bgColor={variables.accentPurple1} tooltip="Footfall" />
 * )
 **/

const Icon = ({ variant, bgColor, className, onClick }: IIconProps) => {
  return (
    <div
      className={`${styles.root} ${styles[variant]} ${className || ""}`}
      style={{ color: bgColor }}
      role="img"
      onClick={onClick}
    >
      {variant && <SvgIcon variant={variant} />}
    </div>
  );
};

const SvgIcon = ({ variant }: { variant: IconVariant }) => {
  switch (variant) {
    case "logo":
      return <LogoIcon className={styles.logo} />;
    case "info":
      return <AiOutlineInfoCircle className={styles.icon} />;
    case "warning":
      return <AiOutlineWarning className={styles.icon} />;
    case "error":
      return <AiOutlineWarning className={styles.icon} />;
    case "success":
      return <HiOutlineCheckCircle className={styles.icon} />;
    case "cart":
      return <AiOutlineShoppingCart className={styles.icon} />;
    case "wallet":
      return <BiWallet className={styles.icon} />;
    case "search":
      return <AiOutlineSearch className={styles.search} />;
    case "ether":
      return <FaEthereum className={styles.icon} />;
    case "upload":
      return <AiOutlineUpload className={styles.icon} />;
    case "setting":
      return <AiOutlineSetting className={styles.icon} />;
    case "book":
      return <GiBookCover className={styles.icon} />;
    case "link":
      return <HiOutlineExternalLink className={styles.icon} />;
    case "exchange":
      return <FaExchangeAlt className={styles.icon} />;
    case "tick":
      return <BsCheck2 className={styles.icon} />;
    case "cross":
      return <RxCross2 className={styles.icon} />;
    case "arrowLeft":
      return <BsArrowLeft className={styles.icon} />;
    case "arrowRight":
      return <BsArrowRight className={styles.icon} />;
    case "filter":
      return <BiFilterAlt className={styles.icon} />;
    case "control":
      return <AiOutlineControl className={styles.icon} />;
    case "gas":
      return <BiGasPump className={styles.icon} />;
    default:
      return null as never;
  }
};

export default Icon;
