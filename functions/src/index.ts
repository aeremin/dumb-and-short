import {setGlobalOptions} from "firebase-functions";
import {onRequest} from "firebase-functions/https";
import * as logger from "firebase-functions/logger";
import * as admin from "firebase-admin";
import {getFirestore} from "firebase-admin/firestore";
import type {Request, Response} from "express";

setGlobalOptions({maxInstances: 10});

admin.initializeApp();

const db = getFirestore();
const urlsCollection = db.collection("urls");

interface MetadataDocument {
  next_id: number;
}

interface UrlDocument {
  url: string;
}

interface CreateRequestBody {
  url: string;
}

interface CreateResponseBody {
  id: string;
}

interface ResolveRequestBody {
  id: string;
}

interface ResolveResponseBody {
  url: string;
}

const handleCors = (req: Request, res: Response): boolean => {
  if (req.method === "OPTIONS") {
    res.header("Access-Control-Allow-Origin", "*");
    res.header("Access-Control-Allow-Methods", "POST");
    res.header("Access-Control-Allow-Headers", "Content-Type");
    res.header("Access-Control-Max-Age", "3600");
    res.sendStatus(204);
    return true;
  }

  res.header("Access-Control-Allow-Origin", "*");
  return false;
};

export const create = onRequest({region: "europe-west6"}, async (req, res) => {
  if (handleCors(req, res)) {
    return;
  }

  if (req.method !== "POST") {
    res.status(405).send("Method not allowed");
    return;
  }

  const body = req.body as CreateRequestBody;
  if (!body || typeof body.url !== "string") {
    res.status(400).send("Can't decode request: invalid body");
    return;
  }

  const responseBody: CreateResponseBody = {id: ""};

  try {
    await db.runTransaction(async (tx) => {
      const metadataRef = urlsCollection.doc("metadata");
      const metaSnap = await tx.get(metadataRef);

      if (!metaSnap.exists) {
        throw new Error("metadata document does not exist");
      }

      const metadata = metaSnap.data() as MetadataDocument;

      const currentId = metadata.next_id;
      const urlDocRef = urlsCollection.doc(String(currentId));

      const urlDoc: UrlDocument = {
        url: body.url,
      };

      tx.create(urlDocRef, urlDoc);

      metadata.next_id = currentId + 1;
      tx.set(metadataRef, metadata);

      responseBody.id = String(currentId);
    });
  } catch (err) {
    logger.error("Internal error in create", err as Error);
    res.status(500).send(`Internal error: ${String(err)}`);
    return;
  }

  res
    .status(200)
    .header("Content-Type", "application/json")
    .json(responseBody);
});

export const resolve = onRequest({region: "europe-west6"}, async (req, res) => {
  if (handleCors(req, res)) {
    return;
  }

  if (req.method !== "POST") {
    res.status(405).send("Method not allowed");
    return;
  }

  const body = req.body as ResolveRequestBody;
  if (!body || typeof body.id !== "string") {
    res.status(400).send("Can't decode request: invalid body");
    return;
  }

  let urlDoc: UrlDocument;

  try {
    const snap = await urlsCollection.doc(body.id).get();
    if (!snap.exists) {
      res.status(404).send("Can't find a document");
      return;
    }

    urlDoc = snap.data() as UrlDocument;
  } catch (err) {
    logger.error("Error reading document in resolve", err as Error);
    res.status(500).send(`Internal error: ${String(err)}`);
    return;
  }

  const responseBody: ResolveResponseBody = {
    url: urlDoc.url,
  };

  res
    .status(200)
    .header("Content-Type", "application/json")
    .json(responseBody);
});
