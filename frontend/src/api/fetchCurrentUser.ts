const apiUrl = import.meta.env.VITE_API_URL;

export const fetchCurrentUser = async (token: string) => {
  try {
    const response = await fetch(`${apiUrl}/api/user/`, {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (!response.ok) {
      throw new Error("Failed to fetch current user");
    }

    const data = await response.json();
    return data.data;
  } catch (error) {
    console.error("Error fetching current user:", error);
    throw error;
  }
};
