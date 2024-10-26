import 'bootstrap/dist/css/bootstrap.min.css';

import Button from 'react-bootstrap/Button';
import Card from 'react-bootstrap/Card';
import InputGroup from 'react-bootstrap/InputGroup';
import FormControl from 'react-bootstrap/FormControl';

import {useState} from "react";

import {create} from "@/api"
import {NextRouter, useRouter} from "next/router";

async function createAndRedirect(router: NextRouter, url: string) {
  const id = await create(url);
  await router.push(`/created/${id}`);
}

export default function Home() {
  const [url, setUrl] = useState("");
  const router = useRouter()

  return (
    <div>
      <Card body={false}>
        <Card.Header>Create a short URL</Card.Header>
        <Card.Body>
          <InputGroup>
            <FormControl type="string" value={url} onChange={(e) => setUrl(e.target.value)} />
            <InputGroup>
              <Button variant="success" onClick={() => createAndRedirect(router, url)}>
                Create
              </Button>
            </InputGroup>
          </InputGroup>
        </Card.Body>
      </Card>
    </div>
  );
}
