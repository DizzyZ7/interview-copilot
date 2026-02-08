import { Link } from "react-router-dom";
import { useAuth } from "../auth/AuthContext";

export function Navbar() {
  const { token, logout } = useAuth();

  return (
    <nav
      style={{
        display: "flex",
        gap: 16,
        padding: 16,
        borderBottom: "1px solid #ddd",
      }}
    >
      {token && (
        <>
          <Link to="/">Dashboard</Link>
          <Link to="/questions">Questions</Link>
          <Link to="/quiz">Quiz</Link>
          <Link to="/progress">Progress</Link>
          <button onClick={logout}>Logout</button>
        </>
      )}
      {!token && (
        <>
          <Link to="/login">Login</Link>
          <Link to="/register">Register</Link>
        </>
      )}
    </nav>
  );
}
