import { useState } from "react";
import { api } from "../api/client";

type Question = {
  id: number;
  title: string;
  body: string;
};

export default function Quiz() {
  const [qs, setQs] = useState<Question[]>([]);
  const [idx, setIdx] = useState(0);

  async function start() {
    const res = await api.post("/api/quiz/start");
    setQs(res.data);
    setIdx(0);
  }

  async function answer(correct: boolean) {
    await api.post("/api/quiz/answer", { correct });
    setIdx(i => i + 1);
  }

  const q = qs[idx];

  return (
    <div>
      <h2>Quiz</h2>
      {qs.length === 0 && <button onClick={start}>Start quiz</button>}
      {q && (
        <div>
          <h3>{q.title}</h3>
          <p>{q.body}</p>
          <button onClick={() => answer(true)}>Correct</button>
          <button onClick={() => answer(false)}>Wrong</button>
        </div>
      )}
      {qs.length > 0 && idx >= qs.length && <p>Done!</p>}
    </div>
  );
}
