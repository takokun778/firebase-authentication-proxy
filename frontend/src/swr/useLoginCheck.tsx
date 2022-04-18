import useSWR from "swr";

import { postLoginCheck } from "../api/api";
import { retryable } from "./retry";

export const useLoginCheck = () => {
    const { data, error } = useSWR(
        "check",
        postLoginCheck,
        {
            onErrorRetry: (err, _key, _config, _revalidate, { retryCount }) => {
                if (!retryable(err, retryCount)) {
                    return;
                }
            },
            revalidateOnFocus: false,
        },
    );

    return { data, error };
};
