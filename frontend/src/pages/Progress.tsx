import { useEffect, useState } from "react";
import { api } from "../api/client";

export default function Progress() {
  const [stats, setStats] = useState<{ total: number; correct: number; ratio: number }>();

  useEffect(() => {
    api.get("/api/progress").then(res => setStats(res.data));
  }, []);

  if (!stats) return <p>Loading...</p>;

  return (
    <div>
      <h2>Progress</h2>
      <p>Total answered: {stats.total}</p>
      <p>Correct: {stats.correct}</p>
      <p>Accuracy: {(stats.ratio * 100).toFixed(1)}%</p>
    </div>
  );
}
