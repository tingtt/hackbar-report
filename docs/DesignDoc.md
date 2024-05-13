# Dssign Doc: hackbar-report

## Objective

HACK.BAR バーテンダーが行うオープン前報告・クローズ報告の作成を効率化する。

## Goal

- インタラクティブな報告書作成コマンド
  - `オープン前報告`
  - `クローズ報告`
  - マークダウン形式で出力

## High Level Structure

- cmd (entrypoint)
  - report
- internal
  - infrastructure
    - cmd
      - report
  - interface-adapter
    - markdown
  - usecase
    - open
    - close
    - prompt-group
  - domain
    - prompt
