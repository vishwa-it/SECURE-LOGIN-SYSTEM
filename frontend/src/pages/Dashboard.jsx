import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import { API_BASE_URL } from "../api";

function Dashboard() {
  const [message, setMessage] = useState("");
  const [role, setRole] = useState("");
  const [nginxUrl, setNginxUrl] = useState("");
  const [error, setError] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const loadDashboard = async () => {
      const token = localStorage.getItem("token");

      if (!token) {
        navigate("/");
        return;
      }

      try {
        const response = await fetch(`${API_BASE_URL}/api/dashboard`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        const data = await response.json();

        if (!response.ok) {
          throw new Error(data.message || "Access denied");
        }

        setMessage(data.message);
        setRole(data.role);
        setNginxUrl(data.nginx_url || "");
      } catch (error) {
        localStorage.removeItem("token");
        setError(error.message || "Unable to load dashboard");
      }
    };

    loadDashboard();
  }, [navigate]);

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/");
  };

  if (error) {
    return (
      <main className="dashboard-container">
        <section className="dashboard-card">
          <h1>Access Denied</h1>
          <p className="dashboard-message">{error}</p>
          <button className="logout-button" onClick={handleLogout}>
            Back to Login
          </button>
        </section>
      </main>
    );
  }

  return (
    <main className="dashboard-container">
      <section className="dashboard-card">
        <div className="dashboard-header">
          <div>
            <p className="eyebrow">Secure Authentication System</p>
            <h1>{role === "admin" ? "Admin Dashboard" : "User Dashboard"}</h1>
            <p className="dashboard-message">{message}</p>
          </div>
          <span className="role-badge">Role: {role}</span>
        </div>

        {role === "admin" && nginxUrl && (
          <div className="nginx-panel">
            <div>
              <h2>Nginx Container</h2>
              <p>Nginx is running through the Docker Engine.</p>
            </div>
            <a href={nginxUrl} target="_blank" rel="noreferrer">
              Open Nginx
            </a>
          </div>
        )}

        {role !== "admin" && (
          <div className="dashboard-item">
            <h2>User Access</h2>
            <p>You are authenticated as a normal user.</p>
          </div>
        )}

        <button className="logout-button" onClick={handleLogout}>
          Logout
        </button>
      </section>
    </main>
  );
}

export default Dashboard;
