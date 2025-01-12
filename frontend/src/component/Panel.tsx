import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { useGetAPIKey } from "../hook/useGetAPIKey";
import APIKeyDisplay from "./APIKeyDisplay";
import { LinearProgress } from "@mui/material";

interface PanelProps {
  username: string;
  token: string;
}

const Panel: React.FC<PanelProps> = ({ username, token }) => {
  const { apikey, expiredAt, loading } = useGetAPIKey(token);

  return (
    <Card
      elevation={24}
      sx={{
        maxWidth: 480,
        p: 3,
        m: 3,
      }}
    >
      <CardContent>
        <Typography variant="h5">Welcome Back,</Typography>
        <Typography variant="h4" fontWeight="bold" sx={{ mb: 1 }}>
          {username}
        </Typography>

        {loading ? (
          <LinearProgress sx={{ mt: 4, mb: 2 }} />
        ) : (
          <>
            <Typography variant="h6" sx={{ mb: 1 }}>
              Here is your API Key:
            </Typography>
            <APIKeyDisplay apiKey={apikey} />
            <Typography variant="body1">
              Your API key will expire on{" "}
              <Typography component="span" style={{ fontWeight: "bold" }}>
                {expiredAt}
              </Typography>
              . Please sign in after expiration to renew it.
            </Typography>
          </>
        )}
      </CardContent>

      <CardActions sx={{ gap: 1 }}>
        <Button variant="contained" size="large">
          <Typography fontWeight={500}>Sign Out</Typography>
        </Button>
        <Button
          href="https://github.com/oadultradeepfield/galaxy10-anomaly-detection?tab=readme-ov-file#updates"
          target="_blank"
          variant="outlined"
          size="large"
        >
          <Typography fontWeight={500}>Learn More</Typography>
        </Button>
      </CardActions>
    </Card>
  );
};

export default Panel;
