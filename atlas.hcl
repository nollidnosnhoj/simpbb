env "local" {
    src = "file://schema.sql"
    dev = "sqlite://simpbb.db"
    migration {
        dir = "file://internal/migrations"
        format = "golang-migrate"
    }
    format {
        migrate {
            diff = "{{ sql . \"  \" }}"
        }
    }
}