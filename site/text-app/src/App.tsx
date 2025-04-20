import { useState } from 'react'
import { DefaultApi, QueryPostRequest } from './api'
import './App.css'

const api = new DefaultApi();

function App() {
  const [results, setResults] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);

  const handleQuery = async () => {
    console.log('Querying...');
    setLoading(true);
    try {
      console.log('Running query...');
      const params: QueryPostRequest = {
        query: 'banana'
      };
      const { data } = await api.queryPost({ queryPostRequest: params });
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
