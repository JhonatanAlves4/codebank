import { MigrationInterface, QueryRunner, Table } from 'typeorm';

export class createProductsTable1675621422796 implements MigrationInterface {
  public async up(queryRunner: QueryRunner): Promise<void> {
    // executar algo no banco: criar tabela, criar campos
    await queryRunner.createTable(
      new Table({
        name: 'products',
        columns: [
          {
            name: 'id',
            type: 'uuid', // unique identifier
            isPrimary: true,
          },
          {
            name: 'name',
            type: 'varchar',
          },
          {
            name: 'description',
            type: 'text',
          },
          {
            name: 'image_url',
            type: 'varchar',
          },
          {
            name: 'slug',
            type: 'varchar',
          },
          {
            name: 'price',
            type: 'double precision',
          },
          {
            name: 'created_at',
            type: 'timestamp',
            default: 'CURRENT_TIMESTAMP',
          },
        ],
      }),
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    // desfazer o que foi feito no up
  }
}
