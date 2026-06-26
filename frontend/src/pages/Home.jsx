import { Link } from 'react-router-dom';
import { useAuth } from '../contexts/AuthContext';
import '../styles/Home.css';

export default function Home() {
  const { isAuthenticated } = useAuth();

  return (
    <div className="home-container">
      <div className="hero">
        <h1>🌍 Rhoam Together</h1>
        <p>Plan your next adventure with friends</p>

        {isAuthenticated ? (
          <Link to="/dashboard" className="btn-primary">
            Go to Dashboard
          </Link>
        ) : (
          <div className="hero-buttons">
            <Link to="/login" className="btn-primary">
              Login
            </Link>
            <Link to="/signup" className="btn-secondary">
              Sign Up
            </Link>
          </div>
        )}
      </div>

      <div className="features">
        <div className="feature">
          <h3>📍 Plan Trips</h3>
          <p>Create and organize your travel itineraries</p>
        </div>
        <div className="feature">
          <h3>👥 Invite Friends</h3>
          <p>Share trips and collaborate with your friends</p>
        </div>
        <div className="feature">
          <h3>💡 Share Ideas</h3>
          <p>Add suggestions and find the best activities</p>
        </div>
      </div>
    </div>
  );
}
