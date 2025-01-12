import { useState } from "react";
import { Box, IconButton } from "@mui/material";
import ContentCopyIcon from "@mui/icons-material/ContentCopy";
import CheckIcon from "@mui/icons-material/Check";
import VisibilityIcon from "@mui/icons-material/Visibility";
import VisibilityOffIcon from "@mui/icons-material/VisibilityOff";

interface APIKeyDisplayProps {
  apiKey: string;
}

const APIKeyDisplay: React.FC<APIKeyDisplayProps> = ({ apiKey }) => {
  const [show, setShow] = useState(false);
  const [copied, setCopied] = useState(false);

  const handleCopy = async () => {
    await navigator.clipboard.writeText(apiKey);
    setCopied(true);
    setTimeout(() => setCopied(false), 1000);
  };

  return (
    <Box
      sx={{
        display: "flex",
        alignItems: "center",
        width: "100%",
        bgcolor: "grey.100",
        p: 1,
        gap: 1,
        my: 2,
      }}
    >
      <Box
        sx={{
          flexGrow: 1,
          fontFamily: "monospace",
          fontSize: "0.875rem",
          overflow: "auto",
          textOverflow: "ellipsis",
          whiteSpace: "nowrap",
        }}
      >
        {show ? apiKey : "•".repeat(apiKey.length / 2)}
      </Box>
      <Box sx={{ display: "flex", gap: 0.5 }}>
        <IconButton
          size="small"
          onClick={handleCopy}
          aria-label={copied ? "Copied" : "Copy to clipboard"}
          sx={{ color: "text.secondary" }}
        >
          {copied ? (
            <CheckIcon fontSize="small" />
          ) : (
            <ContentCopyIcon fontSize="small" />
          )}
        </IconButton>
        <IconButton
          size="small"
          onClick={() => setShow(!show)}
          aria-label={show ? "Hide API key" : "Show API key"}
          sx={{ color: "text.secondary" }}
        >
          {show ? (
            <VisibilityOffIcon fontSize="small" />
          ) : (
            <VisibilityIcon fontSize="small" />
          )}
        </IconButton>
      </Box>
    </Box>
  );
};

export default APIKeyDisplay;
