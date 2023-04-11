import { FC, useState, useRef, CSSProperties } from "react";

import styles from "./imageRenderer.module.scss";
import useIntersection from "hooks/useIntersection";

interface IImageRendererProps {
  url: string;
  style?: CSSProperties;
  className?: string;
}

/**
 * reference: https://dev.to/shubhamreacts/progressively-loading-images-in-react-40lg
 * @param url {string} img link
 * @param className {string} injectable className for positioning
 * @returns ImageRenderer
 *
 * @component
 * @example
 * <ImageRenderer url="" />
 */

const ImageRenderer: FC<IImageRendererProps> = ({ url, style, className }) => {
  const [isInView, setIsInView] = useState(false);
  const imgRef = useRef<HTMLDivElement | null>(null);
  useIntersection(imgRef, () => {
    setIsInView(true);
  });

  return (
    <div className={`${styles.root} ${className || ""}`} ref={imgRef} style={style}>
      {isInView && <img alt="img" className={styles["img"]} src={url} />}
    </div>
  );
};

export default ImageRenderer;
