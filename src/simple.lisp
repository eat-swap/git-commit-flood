(setq cnt %d)
(loop
    ; this file cannot be executed directly.
    (setq cnt (- cnt 1))
    (write-line "%s")
    (when (<= cnt 0)
        (return)
    )
)