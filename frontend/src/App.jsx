import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import './App.css'

// Placeholder pages - will be implemented in Phase 5
function Home() {
  return <h1>Rhoam Together - Welcome</h1>
}

function Login() {
  return <h1>Login (Phase 4)</h1>
}

function Signup() {
  return <h1>Signup (Phase 4)</h1>
}

function Dashboard() {
  return <h1>Dashboard (Phase 5)</h1>
}

export default function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  )
}
