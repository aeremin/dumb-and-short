'use client'

import 'bootstrap/dist/css/bootstrap.min.css';
import Card from "react-bootstrap/Card";

export function SuccessCard({ id } : {id: string}) {
    return (<div>
        <Card body={false}>
            <Card.Header>Success!</Card.Header>
            <Card.Body>
                Link #{id} was successfully created!
            </Card.Body>
        </Card>
    </div>)
}
