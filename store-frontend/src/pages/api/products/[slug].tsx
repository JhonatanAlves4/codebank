// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import { Product, products } from '@/model'
import type { NextApiRequest, NextApiResponse } from 'next'

export default function handler(
  req: NextApiRequest,
  res: NextApiResponse<Product>
) {
  const {slug} = req.query;
  const product = products.find(p => p.slug === slug);
  res.status(200).json(product!)
}
