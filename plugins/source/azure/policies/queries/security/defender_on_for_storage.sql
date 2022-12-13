insert into azure_policy_results
SELECT
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure that Azure Defender is set to On for Storage (Automatic)' as title,
  subscription_id,
  id,
  case
    when pricing_tier = 'Standard'
    then 'pass' else 'fail'
  end
FROM azure_security_pricings asp
WHERE "name" = 'StorageAccounts'
