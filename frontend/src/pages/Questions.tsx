import { useEffect, useState } from "react";
import { api } from "../api/client";

type Question = {
  id: number;
  title: string;
  body: string;
  difficulty: number;
  tags: string[];
};

export default function Questions() {
  const [qs, setQs] = useState<Question[]>([]);
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");

  async function load() {
    const res = await api.get("/api/questions");
    setQs(res.data);
  }

  async function create() {
    await api.post("/api/questions", {
      title,
      body,
      difficulty: 2,
      tags: ["general"],
    });
    setTitle("");
    setBody("");
    load();
  }

  useEffect(() => {
    load();
  }, []);

  return (
    <div>
      <h2>Questions</h2>

      <h3>Create</h3>
      <input placeholder="Title" value={title} onChange={e => setTitle(e.target.value)} />
      <br />
      <textarea placeholder="Body" value={body} onChange={e => setBody(e.target.value)} />
      <br />
      <button onClick={create}>Add</button>

      <h3>List</h3>
      <ul>
        {qs.map(q => (
          <li key={q.id}>
            <b>{q.title}</b> — {q.body}
          </li>
        ))}
      </ul>
    </div>
  );
}
