import { useState, useEffect } from "react";
import { fetchAPIKey } from "../api/fetchAPIKey";

interface APIKeyState {
  apikey: string;
  expiredAt: string;
  loading: boolean;
}

export const useGetAPIKey = (token: string) => {
  const [state, setState] = useState<APIKeyState>({
    apikey: "",
    expiredAt: "",
    loading: false,
  });

  const formatDate = (isoString: string): string => {
    const date = new Date(isoString);
    return new Intl.DateTimeFormat("en-US", {
      year: "numeric",
      month: "long",
      day: "2-digit",
      hour: "2-digit",
      minute: "2-digit",
      second: "2-digit",
      hour12: false,
    }).format(date);
  };

  useEffect(() => {
    const getAPIKey = async () => {
      try {
        const apiKey = await fetchAPIKey(token);
        setState((prev) => ({
          ...prev,
          apikey: apiKey.api_key,
          expiredAt: formatDate(apiKey.expired_at),
          loading: false,
        }));
      } catch (error) {
        setState((prev) => ({
          ...prev,
          apikey: "",
          expiredAt: "",
          loading: false,
        }));
        console.error(error);
      }
    };

    getAPIKey();
  }, [token]);

  return state;
};
