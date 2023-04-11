import { utils } from "ethers/lib/ethers";

const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "numeric",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
    second: "numeric",
    hour12: false
};

export const ethToWei = (amount: number) => {
    return utils.parseEther(amount.toString());
}

export const weiToEth = (amount: string) => {
    if (amount === "undefined") return "";
    return utils.formatEther(amount);
}

export const formatDate = (date: Date) => {
    return new Intl.DateTimeFormat("default", options).format(date);
}
