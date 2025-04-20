import { useState } from 'react'
import { DefaultApi } from './api'
import './App.css'

import { Configuration } from './api/configuration';

const config = new Configuration({ basePath: 'http://localhost:8080/api/v1' });
const api = new DefaultApi(config);

function App() {
  const [results, setResults] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);

  const handleQuery = async () => {
    setLoading(true);
    try {
      const { data } = await api.queryPost({ query: 'banana' });
      setResults(data.results || []);
    } catch (error) {
      console.error('Query failed', error);
    } finally {
      setLoading(false);
    }
  };


  return (
    <>
      <div style={{ padding: '2rem' }}>
        <h1>Search</h1>
        <button onClick={handleQuery} disabled={loading}>
          {loading ? 'Querying...' : 'Run Query'}
        </button>

        <div style={{ marginTop: '1rem' }}>
          <h2>Results:</h2>
          <ul>
            {results.map((snippet, index) => (
              <li key={index}>{snippet}</li>
            ))}
          </ul>
        </div>
      </div>
    </>
  )
}

export default App
