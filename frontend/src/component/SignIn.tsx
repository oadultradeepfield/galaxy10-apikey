import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { Link } from "@mui/material";

const apiUrl = import.meta.env.VITE_API_URL;
const googleOauthUrl = `${apiUrl}/api/auth/google/signin`;

const SignIn = () => {
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
        <Typography variant="h3" fontWeight="bold" sx={{ mb: 1 }}>
          API Key Generator
        </Typography>
        <Typography variant="h5" sx={{ mb: 1 }}>
          Galaxy10 Anomaly Detection
        </Typography>
        <Typography sx={{ color: "text.secondary", mb: 2 }}>
          By{" "}
          <Link
            href="https://github.com/oadultradeepfield"
            underline="hover"
            target="_blank"
            rel="noopener"
          >
            @oadultradeepfield
          </Link>
        </Typography>
        <Typography variant="body1">
          Generate an API key to detect anomaly galaxy images with ResNet50 and
          Autoencoders trained on the Galaxy10 DECals Dataset.
        </Typography>
      </CardContent>
      <CardActions>
        <Button href={googleOauthUrl} variant="contained">
          <Typography variant="body2" fontWeight={500}>
            Get Started
          </Typography>
        </Button>
        <Button
          href="https://github.com/oadultradeepfield/galaxy10-anomaly-detection?tab=readme-ov-file#updates"
          target="_blank"
          variant="outlined"
        >
          <Typography variant="body2" fontWeight={500}>
            Learn More
          </Typography>
        </Button>
      </CardActions>
    </Card>
  );
};

export default SignIn;
