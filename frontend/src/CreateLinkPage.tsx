import 'bootstrap/dist/css/bootstrap.min.css';

import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import InputGroup from 'react-bootstrap/InputGroup';
import FormControl from 'react-bootstrap/FormControl';

import {useState} from "react";
import {create} from "./api";
import {NavigateFunction, useNavigate} from "react-router-dom";

async function createAndRedirect(navigate: NavigateFunction, url: string) {
  const id = await create(url);
  navigate(`/created/${id}`)
}

export default function CreateLinkPage() {
  const [url, setUrl] = useState("");
  const navigate = useNavigate();

  return (
    <div>
      <Card body={false}>
        <Card.Header>Create a short URL</Card.Header>
        <Card.Body>
          <InputGroup>
            <FormControl type="string" value={url} onChange={(e) => setUrl(e.target.value)} />
            <InputGroup>
              <Button variant="success" onClick={() => createAndRedirect(navigate, url)}>
                Create
              </Button>
            </InputGroup>
          </InputGroup>
        </Card.Body>
      </Card>
    </div>
  );
}
