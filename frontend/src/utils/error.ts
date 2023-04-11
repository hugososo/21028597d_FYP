import { INotification } from "features/notificaiton/providers/NotificationProvider";

export const catchError = (err: any) => {
    if (err instanceof Error) {
        return {
            type: "error",
            title: "Something went wrong!",
            content: err.message,
        } satisfies INotification;
    } else {
        return {
            type: "error",
            title: "Something went wrong!",
            content: (err as any).data.message,
        } satisfies INotification;
    }
} 