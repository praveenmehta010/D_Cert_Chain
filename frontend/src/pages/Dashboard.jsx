import { useState } from "react";
import axios from "axios";
import Navbar from "../components/navbar";

function Dashboard() {
  const [mode, setMode] = useState("");
  const [issueForm, setIssueForm] = useState({
    studentName: "",
    studentEmail: "",
    course: "",
    issuerName: "",
    noOfIssue: "",
  });
  const [revokeForm, setRevokeForm] = useState({ certificateHash: "" });
  const [message, setMessage] = useState("");

  const handleIssue = async () => {
    try {
      console.log("📦 Form data being sent:", issueForm);

      const res = await axios.post("http://localhost:8080/issue", issueForm, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
      setMessage("✅ Certificate Issued");
    } catch (err) {
      setMessage("❌ Failed to Issue");
    }
  };

  const handleRevoke = async () => {
    try {
      const res = await axios.post("http://localhost:8080/revoke", revokeForm, {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      });
      setMessage("✅ Certificate Revoked");
    } catch (err) {
      setMessage("❌ Failed to Revoke");
    }
  };

  return (
    <div>
      <Navbar showLogout={true} />
      <div className="flex justify-center items-center h-screen">
        <div className="w-full max-w-xl bg-amber-200 p-7 rounded-lg shadow-md">
          <h1 className="text-2xl font-bold text-center mb-4 bg-blue-400 p-2 rounded">
            🎓 Certificate Actions
          </h1>

          {/* Action Toggle Buttons */}
          <div className="flex justify-center space-x-4 mb-6">
            <button
              className={`px-4 py-2 rounded ${mode === "issue" ? "bg-blue-600 text-white" : "bg-white border"}`}
              onClick={() => {
                setMode("issue");
                setMessage("");
              }}
            >
              Issue Certificate
            </button>
            <button
              className={`px-4 py-2 rounded ${mode === "revoke" ? "bg-blue-600 text-white" : "bg-white border"}`}
              onClick={() => {
                setMode("revoke");
                setMessage("");
              }}
            >
              Revoke Certificate
            </button>
          </div>

          {/* Conditional Form */}
          {mode === "issue" && (
            <div className="mb-6">
              <h2 className="font-semibold mb-2">📝 Issue Certificate</h2>
              <input type="text" placeholder="Name" className="border p-2 w-full mb-2" onChange={(e) => setIssueForm({ ...issueForm, studentName: e.target.value })} />
              <input type="email" placeholder="Email" className="border p-2 w-full mb-2" onChange={(e) => setIssueForm({ ...issueForm, studentEmail: e.target.value })} />
              <input type="text" placeholder="Course" className="border p-2 w-full mb-2" onChange={(e) => setIssueForm({ ...issueForm, course: e.target.value })} />
              <input type="text" placeholder="Issuer Organization" className="border p-2 w-full mb-2" onChange={(e) => setIssueForm({ ...issueForm, issuerName: e.target.value })} />
              <input type="text" placeholder="Issue no" className="border p-2 w-full mb-2" onChange={(e) => setIssueForm({ ...issueForm, noOfIssue: e.target.value })} />
              <button className="bg-green-600 text-white px-4 py-2 rounded" onClick={handleIssue}>
                Issue
              </button>
            </div>
          )}

          {mode === "revoke" && (
            <div>
              <h2 className="font-semibold mb-2">⛔ Revoke Certificate</h2>
              <input type="text" placeholder="Certificate ID" className="border p-2 w-full mb-2" onChange={(e) => setRevokeForm({ certificateHash: e.target.value })} />
              <button className="bg-red-600 text-white px-4 py-2 rounded" onClick={handleRevoke}>
                Revoke
              </button>
            </div>
          )}

          {message && (
            <div className="mt-4 text-center font-semibold">
              {message}
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default Dashboard;