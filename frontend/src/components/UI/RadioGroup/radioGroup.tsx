import { FC } from "react";

import styles from "./radioGroup.module.scss";

export interface IRadioOption {
  text: string;
  value: string | number | readonly string[] | undefined;
  color?: "red" | "green";
}

interface IRadioGroupProps {
  groupName: string;
  labelValueList: IRadioOption[];
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void;
  horizontal?: boolean;
  selected?: string;
}

const RadioGroup: FC<IRadioGroupProps> = ({ groupName, labelValueList, onChange, horizontal, selected }) => {
  return (
    <div className={horizontal ? styles["radio-group-horizontal"] : styles["radio-group-vertical"]}>
      {labelValueList.map((i) => {
        return (
          <label key={i.text} className={`${styles["radio"]} ${horizontal ? styles["radio-margin-right"] : ""}`}>
            <input
              type="radio"
              color={i.color}
              value={i.value}
              name={groupName}
              onChange={onChange}
              checked={i.value === selected}
            />
            {i.text}
          </label>
        );
      })}
    </div>
  );
};

export default RadioGroup;
