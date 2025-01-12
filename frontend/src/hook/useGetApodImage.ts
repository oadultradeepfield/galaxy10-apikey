import { useState, useEffect } from "react";
import { fetchApodImage } from "../api/fetchApodImage";

interface ApodImageState {
  backgroundImage: string;
  loading: boolean;
}

export const useGetApodImage = () => {
  const [state, setState] = useState<ApodImageState>({
    backgroundImage: "",
    loading: true,
  });

  useEffect(() => {
    const getApodImage = async () => {
      try {
        const imageUrl = await fetchApodImage();
        setState((prev) => ({
          ...prev,
          backgroundImage: imageUrl,
          loading: false,
        }));
      } catch (error) {
        console.error(error);
      }
    };

    getApodImage();
  }, []);

  return state;
};
