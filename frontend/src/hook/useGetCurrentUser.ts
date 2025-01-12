import { useState, useEffect } from "react";
import { fetchCurrentUser } from "../api/fetchCurrentUser";

interface CurrentUserState {
  username: string;
  isLoaded: boolean;
}

export const useGetCurrentUser = (token: string) => {
  const [state, setState] = useState<CurrentUserState>({
    username: "",
    isLoaded: false,
  });

  useEffect(() => {
    const getCurrentUser = async () => {
      try {
        const userInfo = await fetchCurrentUser(token);
        setState((prev) => ({
          ...prev,
          username: userInfo.username,
          isLoaded: true,
        }));
      } catch (error) {
        setState((prev) => ({
          ...prev,
          username: "",
          isLoaded: true,
        }));
        console.error(error);
      }
    };

    getCurrentUser();
  }, []);

  return state;
};
