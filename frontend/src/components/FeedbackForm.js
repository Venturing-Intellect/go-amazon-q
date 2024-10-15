import React, { useState } from 'react';
import axios from 'axios';

function FeedbackForm() {
    const [name, setName] = useState('');
    const [email, setEmail] = useState('');
    const [feedback, setFeedback] = useState('');
    const [message, setMessage] = useState('');
    const [isError, setIsError] = useState(false);
  
    const handleSubmit = async (e) => {
      e.preventDefault();
      try {
        const response = await axios.post('http://localhost:8080/feedback', { name, email, feedback });
        setMessage('Feedback submitted successfully!');
        setIsError(false);
        setName('');
        setEmail('');
        setFeedback('');
      } catch (error) {
        setMessage('Error submitting feedback. Please try again.');
        setIsError(true);
      }
    };
  
    return (
      <div className="row justify-content-center">
        <div className="col-md-6">
          <form onSubmit={handleSubmit}>
            <div className="mb-3">
              <label htmlFor="name" className="form-label">Name</label>
              <input
                type="text"
                className="form-control"
                id="name"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
              />
            </div>
            <div className="mb-3">
              <label htmlFor="email" className="form-label">Email</label>
              <input
                type="email"
                className="form-control"
                id="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
              />
            </div>
            <div className="mb-3">
              <label htmlFor="feedback" className="form-label">Feedback</label>
              <textarea
                className="form-control"
                id="feedback"
                value={feedback}
                onChange={(e) => setFeedback(e.target.value)}
                required
              ></textarea>
            </div>
            <button type="submit" className="btn btn-primary">Submit Feedback</button>
          </form>
          {message && (
            <div className={`alert mt-3 ${isError ? 'alert-danger' : 'alert-success'}`}>
              {message}
            </div>
          )}
        </div>
      </div>
    );
  }
  
export default FeedbackForm;
