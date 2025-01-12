const apiUrl = import.meta.env.VITE_API_URL;

export const fetchAPIKey = async (token: string) => {
  try {
    const response = await fetch(`${apiUrl}/api/apikey/`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Failed to fetch the api key");
    }

    const data = await response.json();
    return data.data;
  } catch (error) {
    console.error("Error fetching api key:", error);
    throw error;
  }
};
