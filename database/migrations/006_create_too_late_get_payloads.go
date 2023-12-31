package migrations

import (
	"github.com/flashbots/mev-boost-relay/database/vars"
	migrate "github.com/rubenv/sql-migrate"
)

var Migration006CreateTooLateGetPayload = &migrate.Migration{
	Id: "006-create-too-late-get-payload",
	Up: []string{`
		CREATE TABLE IF NOT EXISTS ` + vars.TableTooLateGetPayload + ` (
			id          bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
			inserted_at timestamp NOT NULL default current_timestamp,

			slot            bigint NOT NULL,

			slot_start_timestamp   bigint NOT NULL,
			request_timestamp bigint NOT NULL,
			decode_timestamp  bigint NOT NULL,

			proposer_pubkey varchar(98) NOT NULL,
			block_hash      varchar(66) NOT NULL,
			ms_into_slot    bigint NOT NULL
		);

		CREATE UNIQUE INDEX IF NOT EXISTS ` + vars.TableTooLateGetPayload + `_slot_pk_hash_idx ON ` + vars.TableTooLateGetPayload + `(slot, proposer_pubkey, block_hash);
	`, `
		ALTER TABLE ` + vars.TableDeliveredPayload + ` ADD publish_ms bigint NOT NULL DEFAULT 0;
	`},
	Down: []string{},

	DisableTransactionUp:   true,
	DisableTransactionDown: true,
}
