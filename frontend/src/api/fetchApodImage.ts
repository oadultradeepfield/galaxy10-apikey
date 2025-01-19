const nasaApiKey = import.meta.env.VITE_NASA_API_KEY;
const apodApiUrl = `https://api.nasa.gov/planetary/apod?api_key=${nasaApiKey}&date=2006-03-16`;

export const fetchApodImage = async () => {
  try {
    const response = await fetch(apodApiUrl);
    if (!response.ok) {
      throw new Error("Failed to fetch image");
    }
    const data = await response.json();
    return data.hdurl;
  } catch (error) {
    console.error("Error fetching NASA image:", error);
    throw error;
  }
};
