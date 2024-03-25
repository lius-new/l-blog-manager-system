import { Routes, Route, Navigate } from "react-router-dom";
import { Dashboard, Auth, Errors } from "@/layouts";

function App() {
  return (
    <Routes>
      <Route path="/dashboard/*" element={<Dashboard />} />
      <Route path="/auth/*" element={<Auth />} />
      <Route path="/errors/*" element={<Errors />} />
      <Route
        path="/"
        element={<Navigate to="/auth/sign-in" replace={true} />}
      />
      <Route
        path="*"
        element={<Navigate to="/errors/not-found" replace={true} />}
      />
    </Routes>
  );
}

export default App;
