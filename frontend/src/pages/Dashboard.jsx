import { useAuth } from '../contexts/AuthContext';
import '../styles/Dashboard.css';

export default function Dashboard() {
  const { user } = useAuth();

  return (
    <div className="dashboard-container">
      <h1>Welcome, {user?.name}!</h1>

      <div className="dashboard-content">
        <p>This is your dashboard. Coming in Phase 6:</p>
        <ul>
          <li>View your trips</li>
          <li>Create new trips</li>
          <li>Invite friends</li>
          <li>Add suggestions for activities</li>
          <li>Real-time updates</li>
        </ul>
      </div>
    </div>
  );
}
