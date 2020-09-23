import React from 'react';
import { Container, Input, Button } from 'semantic-ui-react';

function App() {
  return (
    <Container>
      <h1>Schedule notification</h1>
      <Input placeholder="Title" />
      <Input placeholder="Message" />
      <Button content="Preview" />
    </Container>
  );
}

export default App;
