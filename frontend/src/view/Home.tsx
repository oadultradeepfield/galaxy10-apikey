import { useState, useEffect } from "react";
import { Box } from "@mui/material";
import SignIn from "../component/SignIn";
import Panel from "../component/Panel";
import Loading from "../component/Loading";
import { useGetApodImage } from "../hook/useGetApodImage";
import { useGetCurrentUser } from "../hook/useGetCurrentUser";

const Home = () => {
  const [token, setToken] = useState(
    localStorage.getItem("galaxy10apiKeyAccessToken") || ""
  );

  useEffect(() => {
    const handleTokenCheck = () => {
      try {
        const urlParams = new URLSearchParams(window.location.search);
        const tokenFromUrl = urlParams.get("token");

        if (tokenFromUrl) {
          localStorage.setItem("galaxy10apiKeyAccessToken", tokenFromUrl);
          urlParams.delete("galaxy10apiKeyAccessToken");
          window.history.replaceState(
            {},
            document.title,
            window.location.pathname
          );

          setToken(tokenFromUrl);
        }
      } catch (error) {
        console.error(error);
      }
    };
    handleTokenCheck();
  }, []);

  const { backgroundImage, loading } = useGetApodImage();
  const { username, isLoaded } = useGetCurrentUser(token);

  if (loading || !isLoaded) {
    return <Loading />;
  }

  return (
    <Box
      sx={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        width: 1,
        height: "100vh",
        backgroundImage: `url(${backgroundImage})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        overflow: "hidden",
        position: "fixed",
        top: 0,
        left: 0,
      }}
    >
      {username ? <Panel username={username} token={token} /> : <SignIn />}
    </Box>
  );
};

export default Home;
